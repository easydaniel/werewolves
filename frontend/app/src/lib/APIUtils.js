const HOSTNAME = "http://10.0.0.20:8081"

// User
export const registerUser = async (username, password) =>
  await fetch(`${HOSTNAME}/users/register`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  }).then(async resp => {
    if (resp.ok) {
      return [{ username }, null]
    }
    const { error } = await resp.json()

    return [null, error]
  })

export const login = async (username, password) =>
  await fetch(`${HOSTNAME}/users/login`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  }).then(async resp => {
    if (resp.ok) {
      return [{ username }, null]
    }
    const { error } = await resp.json()

    return [null, error]
  })

export const logout = async () => {
  await fetch(`${HOSTNAME}/users/logout`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
  })
  location.reload()
}

// Game
export const getBoardList = async () =>
  await fetch(`${HOSTNAME}/boardtype`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  }).then(async resp => {
    if (resp.ok) {
      return await resp.json()
    }
    return null
  })

export const createGame = async board =>
  await fetch(`${HOSTNAME}/games`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ board }),
  }).then(async resp => resp.json())

export const getGameStatus = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  }).then(async resp => resp.json())

export const testGameStart = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}/test/start`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  }).then(async resp => resp.json())
