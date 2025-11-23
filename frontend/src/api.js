export async function api(path, method = "GET", body = null, token = null) {
  const options = {
    method,
    headers: {
      "Content-Type": "application/json",
    },
  };

  if (token) {
    options.headers["Authorization"] = "Bearer " + token;
  }

  if (body) {
    options.body = JSON.stringify(body);
  }

  const res = await fetch("http://localhost:8080" + path, options);

  let data = null;
  try {
    data = await res.json();
  } catch (_) {}

  return { ok: res.ok, status: res.status, data };
}
