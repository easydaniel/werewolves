const HOSTNAME = "http://localhost:8081"

export const registerUser = async (username, password) => {
  const resp = await fetch(`${HOSTNAME}/users/register`, {
    mode: "no-cors",
    method: "POST",
    body: { name: username, password },
  })
  return resp
}

export const login = async (username, password) => {
  const resp = await fetch(`${HOSTNAME}/users/login`, {
    mode: "no-cors",
    method: "POST",
    body: { name: username, password },
  })
  return true // return ?user
}

export const logout = async () => {
  await fetch(`${HOSTNAME}/users/logout`, {
    mode: "no-cors",
    method: "POST",
  })
  location.reload()
}
