import React, { useState } from "react";
import {
  MessageOutlined,
  MenuUnfoldOutlined,
  MenuFoldOutlined,
  PlusOutlined,
} from "@ant-design/icons";
import { Layout, Menu, theme, Button } from "antd";

import Chat from "../OpenAiRef/Chat";

const { Content, Sider, Header } = Layout;
const items = [MessageOutlined].map((icon, index) => ({
  key: String(index + 1),
  icon: React.createElement(icon),
  label: `Чат ${index + 1}`,
}));

const ChatMain = ({ connectionString }) => {
  const [collapsed, setCollapsed] = useState(false);

  const [chats, setChats] = useState([
    {
      key: String(1),
      icon: React.createElement(MessageOutlined),
      label: `Чат ${1}`,
    },
  ]);

  const toggleCollapsed = () => {
    setCollapsed(!collapsed);
  };

  const onAddChat = () => {
    setChats([
      ...chats,
      {
        key: String(chats.length + 1),
        icon: React.createElement(MessageOutlined),
        label: `Чат ${chats.length + 1}`,
      },
    ]);
  };

  return (
    <Layout hasSider className="overflow-hidden">
      <Sider className="h-screen" collapsed={collapsed}>
        <div className="">
          <Button type="primary" onClick={toggleCollapsed}>
            {collapsed ? <MenuUnfoldOutlined /> : <MenuFoldOutlined />}
          </Button>
        </div>
        <Menu
          mode="inline"
          theme="dark"
          items={chats}
          className="h-5/6 overflow-scroll"
        ></Menu>
        <Button
          icon={<PlusOutlined />}
          type="primary"
          block
          className="mt-auto"
          onClick={onAddChat}
        >
          Чат
        </Button>
      </Sider>
      <Layout>
        <Content className="h-screen px-2">
          <Chat connectionString={connectionString} />
        </Content>
      </Layout>
    </Layout>
  );
};
export default ChatMain;
