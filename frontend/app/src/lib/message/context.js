import React from "react"

const MessageContext = React.createContext({
  message: { severity: null, value: null },
  setMessage: () => {},
})

export const MessageProvider = MessageContext.Provider
export default MessageContext
