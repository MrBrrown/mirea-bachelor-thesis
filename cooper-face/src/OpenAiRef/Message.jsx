import React from "react";
import { Avatar, Typography } from "antd";

const { Text } = Typography;

const Message = ({ sender, text, timestamp }) => {
  return (
    <div
      style={{
        display: "flex",
        flexDirection: "column",
        alignItems: "flex-start",
      }}
    >
      <div style={{ display: "flex", alignItems: "center" }}>
        <Avatar>{sender.charAt(0)}</Avatar>
        <div
          style={{ display: "flex", flexDirection: "column", marginLeft: 10 }}
        >
          <Text strong>{sender}</Text>
        </div>
      </div>
      <div style={{ display: "flex", alignItems: "center" }}>
        <Text>{text}</Text>
        <Text type="secondary" style={{ marginLeft: 10 }}>
          {timestamp}
        </Text>
      </div>
    </div>
  );
};

export default Message;
