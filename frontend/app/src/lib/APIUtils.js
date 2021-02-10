const HOSTNAME = "http://10.0.0.20:8081"

const middlewares = {
  errorHandler: async resp => {
    if (resp.ok) {
      return resp.json()
    }
    const { error } = await resp.json()
    return [null, error]
  },
}

// User
export const registerUser = async (username, password) =>
  await fetch(`${HOSTNAME}/users/register`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])

export const login = async (username, password) =>
  await fetch(`${HOSTNAME}/users/login`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ name: username, password }),
  })
    .then(middlewares.errorHandler)
    .then(json => [{ username }, null])

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
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])

export const createGame = async board =>
  await fetch(`${HOSTNAME}/games`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ board }),
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])

export const joinRoom = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])

export const exitRoom = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "DELETE",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])

export const getGameStatus = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])

export const setSeat = async (gameID, seat) =>
  await fetch(`${HOSTNAME}/games/${gameID}/seat`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ seat }),
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])

export const exitSeat = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}/seat`, {
    credentials: "include",
    mode: "cors",
    method: "DELETE",
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])

export const testGameStart = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}/test/start`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
