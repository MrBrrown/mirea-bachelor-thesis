import React, { useState } from "react";
import MessageList from "./MessageList";
import MessageInput from "./MessageInput";

const Chat = ({ connectionString }) => {
  const [messages, setMessages] = useState([]);

  const sendMessage = (text) => {
    if (text[0] === "/") {
      const newMessage = {
        sender: "Cooper",
        text: text.slice(1),
        timestamp: "",
      };
      setMessages([...messages, newMessage]);
    } else {
      const newMessage = {
        sender: "User",
        text: text,
        timestamp: "",
      };
      setMessages([...messages, newMessage]);
    }
  };

  return (
    <>
      <div className="h-5/6 overflow-scroll py-5">
        <MessageList messages={messages} />
      </div>
      <MessageInput
        connectionString={connectionString}
        onSendMessage={sendMessage}
      />
    </>
  );
};
export default Chat;
