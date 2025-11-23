export default async function api(path, method = "GET", body = null) {
  try {
    const token = localStorage.getItem("token");

    const res = await fetch("http://localhost:8080" + path, {
      method,
      headers: {
        "Content-Type": "application/json",
        Authorization: token ? `Bearer ${token}` : "",
      },
      body: body ? JSON.stringify(body) : null,
    });

    const data = await res.json().catch(() => ({}));

    return { ok: res.ok, data };
  } catch (err) {
    alert("Server unreachable");
    return { ok: false, data: null };
  }
}
