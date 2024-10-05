import React, { useState } from "react";
import { Input, Button } from "antd";

import { UpCircleOutlined } from "@ant-design/icons";

import AudioRecorderComponent from "../AudioRecorder/AudioRecorder";

const { TextArea } = Input;

const MessageInput = ({ connectionStinrg, onSendMessage }) => {
  const [message, setMessage] = useState("");
  const [isEmpty, setIsEmpty] = useState(true);

  const checkMessage = (message) => {
    if (message === "") {
      return false;
    } else {
      return true;
    }
  };

  const handleChange = (e) => {
    setMessage(e.target.value);

    if (checkMessage(e.target.value)) {
      if (isEmpty) {
        setIsEmpty(false);
      }
    } else {
      if (!isEmpty) {
        setIsEmpty(true);
      }
    }
  };

  const handleSubmit = () => {
    if (checkMessage(message)) {
      onSendMessage(message);
      setMessage("");
      setIsEmpty(true);
    }
  };

  const handleKeyPress = (e) => {
    if (e.key === "Enter" && !e.shiftKey) {
      e.preventDefault();
      handleSubmit();
    }
  };

  return (
    <>
      <div className="flex gap-1">
        <TextArea
          rows={3}
          value={message}
          onChange={handleChange}
          onKeyDown={handleKeyPress}
          style={{ resize: "none" }}
        />
        <div className="">
          <Button
            disabled={isEmpty}
            icon={<UpCircleOutlined />}
            onClick={handleSubmit}
          ></Button>
        </div>
      </div>
      <div>
        <AudioRecorderComponent
          connectionStinrg={connectionStinrg}
          onServerResponse={onSendMessage}
        ></AudioRecorderComponent>
      </div>
    </>
  );
};

export default MessageInput;
