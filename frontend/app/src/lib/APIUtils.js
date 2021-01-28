const HOSTNAME = "http://localhost:8081"

export const registerUser = async (username, password) => {
  const resp = await fetch(`${HOSTNAME}/users/register`, {
    mode: "no-cors",
    method: "POST",
    body: { name: username, password },
  }).fail(err => console.log(err))
  return resp
}
