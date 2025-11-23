import React, { useEffect, useState } from 'react';
import  api from './api';

export default function Items({ token }) {
  const [items, setItems] = useState([]);
  const [cart, setCart] = useState(null);

  useEffect(()=> {
    loadItems();
  }, []);

  async function loadItems() {
    const res = await api('/items');
    if (res.ok) setItems(res.data || []);
  }

  async function addToCart(itemId) {
    const res = await api('/carts', 'POST', { item_id: itemId }, token);
    if (res.ok) {
      setCart(res.data);
      window.alert('Item added to cart');
    } else {
      window.alert('Failed to add to cart: ' + (res.data?.error || res.status));
    }
  }

  async function viewCart() {
    if (!cart) {
      // try fetch user's active cart by listing carts and filtering
      const res = await api('/carts');
      if (!res.ok) return window.alert('no cart');
      const all = res.data || [];
      const own = all.find(c => c.status === 'active' && c.items && c.items.length>0);
      if (!own) return window.alert('cart empty');
      const ids = (own.items||[]).map(it=>it.item_id);
      return window.alert('Cart items: ' + JSON.stringify(ids));
    }
    const ids = (cart.items || []).map(it => it.item_id);
    window.alert('Cart items: ' + JSON.stringify(ids));
  }

  async function viewOrders() {
    const res = await api('/orders', 'GET', null, token);
    if (!res.ok) return window.alert('Failed to fetch orders');
    const ids = (res.data || []).map(o => o.id);
    window.alert('Order IDs: ' + JSON.stringify(ids));
  }

  async function checkout() {
    // need cart id: if cart present use it, else attempt to find active cart
    let cartId = cart?.id;
    if (!cartId) {
      const res = await api('/carts');
      if (!res.ok) return window.alert('No cart');
      const all = res.data || [];
      const own = all.find(c => c.status === 'active');
      if (!own) return window.alert('no active cart');
      cartId = own.id;
    }
    const res = await api('/orders', 'POST', { cart_id: cartId }, token);
    if (!res.ok) return window.alert('Checkout failed: ' + (res.data?.error || res.status));
    window.alert('Order successful');
    // refresh items/clear local cart state
    setCart(null);
  }

  return (
    <div>
      <div className="topbar">
        <button onClick={checkout}>Checkout</button>
        <button onClick={viewCart}>Cart</button>
        <button onClick={viewOrders}>Order History</button>
      </div>

      <h2>Items</h2>
      <div className="items-list">
        {items.map(item => (
          <div key={item.id} className="item" onClick={()=>addToCart(item.id)}>
            <strong>{item.name}</strong>
            <div>{item.status}</div>
          </div>
        ))}
      </div>
    </div>
  );
}
