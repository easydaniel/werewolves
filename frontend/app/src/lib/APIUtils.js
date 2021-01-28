const HOSTNAME = "http://192.168.137.102:8081"

export const registerUser = async (username, password) => {
  const resp = await fetch(`${HOSTNAME}/users/register`, {
    credentials: 'include',
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  })
  return resp
}

export const login = async (username, password) => {
  const resp = await fetch(`${HOSTNAME}/users/login`, {
    credentials: 'include',
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  })
  return true // return ?user
}

export const logout = async () => {
  await fetch(`${HOSTNAME}/users/logout`, {
    credentials: 'include',
    mode: "cors",
    method: "POST",
  })
  location.reload()
}
