const HOSTNAME = "http://192.168.102.108:8081"

const middlewares = {
  errorHandler: async resp => {
    if (!resp.ok) {
      const { error } = await resp.json()
      throw error
    }
    return resp.json()
  },
}

// User
export const registerUser = async (username, password) =>
  await fetch(`${HOSTNAME}/users/register`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ username, password }),
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])
    .catch(error => [null, error])

export const getUser = async () =>
  await fetch(`${HOSTNAME}/users/profile`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
    .catch(error => [null, error])

export const login = async (username, password) =>
  await fetch(`${HOSTNAME}/users/login`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ username, password }),
  })
    .then(middlewares.errorHandler)
    .then(json => [{ username }, null])
    .catch(error => [null, error])

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
    .catch(error => [null, error])

export const createGame = async board =>
  await fetch(`${HOSTNAME}/games`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ board }),
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
    .catch(error => [null, error])

export const joinRoom = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])
    .catch(error => [null, error])

export const exitRoom = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "DELETE",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
    .catch(error => [null, error])

export const getGameStatus = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
    .catch(error => [null, error])

export const setSeat = async (gameID, seat) =>
  await fetch(`${HOSTNAME}/games/${gameID}/seat`, {
    credentials: "include",
    mode: "cors",
    method: "POST",
    body: JSON.stringify({ seat }),
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])
    .catch(error => [null, error])

export const exitSeat = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}/seat`, {
    credentials: "include",
    mode: "cors",
    method: "DELETE",
  })
    .then(middlewares.errorHandler)
    .then(json => ["Success", null])
    .catch(error => [null, error])

export const testGameStart = async gameID =>
  await fetch(`${HOSTNAME}/games/${gameID}/test/start`, {
    credentials: "include",
    mode: "cors",
    method: "GET",
  })
    .then(middlewares.errorHandler)
    .then(json => [json, null])
    .catch(error => [null, error])
