package main

import (
    "net/http"
    "time"

    "github.com/gin-gonic/gin"
    "github.com/google/uuid"
)

// CreateUser: POST /users
func CreateUser(c *gin.Context) {
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&body); err != nil || body.Username == "" || body.Password == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    user := User{Username: body.Username, Password: body.Password, CreatedAt: time.Now()}
    if err := DB.Create(&user).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "username already exists"})
        return
    }
    c.JSON(http.StatusCreated, gin.H{"id": user.ID, "username": user.Username})
}

// ListUsers: GET /users
func ListUsers(c *gin.Context) {
    var users []User
    DB.Find(&users)
    // hide password
    for i := range users {
        users[i].Password = ""
    }
    c.JSON(http.StatusOK, users)
}

// LoginUser: POST /users/login
func LoginUser(c *gin.Context) {
    var body struct {
        Username string `json:"username"`
        Password string `json:"password"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    var user User
    if err := DB.Where("username = ? AND password = ?", body.Username, body.Password).First(&user).Error; err != nil {
        c.JSON(http.StatusUnauthorized, gin.H{"error": "invalid username/password"})
        return
    }
    // generate a token and save (single-token-per-user)
    token := uuid.New().String()
    user.Token = token
    DB.Save(&user)

    c.JSON(http.StatusOK, gin.H{"token": token})
}

// CreateItem: POST /items
func CreateItem(c *gin.Context) {
    var body struct {
        Name string `json:"name"`
    }
    if err := c.BindJSON(&body); err != nil || body.Name == "" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }
    item := Item{Name: body.Name, Status: "available", CreatedAt: time.Now()}
    DB.Create(&item)
    c.JSON(http.StatusCreated, item)
}

// ListItems: GET /items
func ListItems(c *gin.Context) {
    var items []Item
    DB.Find(&items)
    c.JSON(http.StatusOK, items)
}

// CreateOrAddToCart: POST /carts
// Expect payload: { "item_id": 1 } or { "item_ids":[1,2] }
func CreateOrAddToCart(c *gin.Context) {
    ui, _ := c.Get("user")
    user := ui.(User)

    var payload struct {
        ItemID  *uint  `json:"item_id"`
        ItemIDs []uint `json:"item_ids"`
        Name    string `json:"name"`
    }
    if err := c.BindJSON(&payload); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    // find or create active cart for this user
    var cart Cart
    if err := DB.Where("user_id = ? AND status = ?", user.ID, "active").First(&cart).Error; err != nil {
        // create new
        cart = Cart{UserID: user.ID, Name: payload.Name, Status: "active", CreatedAt: time.Now()}
        DB.Create(&cart)
        // set user's cart_id
        DB.Model(&user).Update("cart_id", cart.ID)
    }

    // add items
    if payload.ItemID != nil {
        ci := CartItem{CartID: cart.ID, ItemID: *payload.ItemID}
        DB.Create(&ci)
    }
    for _, iid := range payload.ItemIDs {
        ci := CartItem{CartID: cart.ID, ItemID: iid}
        DB.Create(&ci)
    }

    // load items to return
    DB.Preload("Items").First(&cart, cart.ID)
    c.JSON(http.StatusOK, cart)
}

// ListCarts: GET /carts
func ListCarts(c *gin.Context) {
    var carts []Cart
    DB.Preload("Items").Find(&carts)
    c.JSON(http.StatusOK, carts)
}

// CreateOrderFromCart: POST /orders
// payload: { "cart_id": 1 }
func CreateOrderFromCart(c *gin.Context) {
    ui, _ := c.Get("user")
    user := ui.(User)

    var body struct {
        CartID uint `json:"cart_id"`
    }
    if err := c.BindJSON(&body); err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "invalid payload"})
        return
    }

    var cart Cart
    if err := DB.Preload("Items").First(&cart, body.CartID).Error; err != nil {
        c.JSON(http.StatusBadRequest, gin.H{"error": "cart not found"})
        return
    }
    if cart.UserID != user.ID {
        c.JSON(http.StatusForbidden, gin.H{"error": "cart does not belong to user"})
        return
    }
    if cart.Status == "ordered" {
        c.JSON(http.StatusBadRequest, gin.H{"error": "cart already ordered"})
        return
    }

    // create order
    order := Order{CartID: cart.ID, UserID: user.ID, CreatedAt: time.Now()}
    DB.Create(&order)

    // mark cart as ordered
    cart.Status = "ordered"
    DB.Save(&cart)

    c.JSON(http.StatusCreated, order)
}

// ListOrders: GET /orders
func ListOrders(c *gin.Context) {
    var orders []Order
    DB.Find(&orders)
    c.JSON(http.StatusOK, orders)
}
