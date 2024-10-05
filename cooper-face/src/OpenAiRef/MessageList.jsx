import React from "react";
import Message from "./Message";

const MessageList = ({ messages }) => {
  return (
    <div>
      {messages.map((message, index) => (
        <Message
          key={index}
          sender={message.sender}
          text={message.text}
          timestamp={message.timestamp}
        />
      ))}
    </div>
  );
};

export default MessageList;
