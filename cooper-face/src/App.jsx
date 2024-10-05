import React, { useState } from "react";

import ChatMain from "./Main/Chat";
import QRScanner from "./QrScanner/TestQR";

const App = () => {
  const [isInit, setIsInit] = useState(false);
  const [connection, setConnection] = useState("");

  const onInitServer = (text) => {
    setIsInit(true);
    setConnection(text);
  };

  return (
    <>
      {!isInit ? <ChatMain connectionString={connection} /> : null}
      {isInit ? <QRScanner onInitServer={onInitServer} /> : null}
    </>
  );
};
export default App;
