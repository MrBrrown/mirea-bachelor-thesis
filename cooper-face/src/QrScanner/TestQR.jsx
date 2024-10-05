import { Scanner } from "@yudiel/react-qr-scanner";
import baseURL from "../../config";
import { useState } from "react";

const QRScanner = ({ onInitServer }) => {
  const [isRequested, setIsRequested] = useState(false);

  const handleScan = (text) => {
    if (isRequested) {
      return;
    }

    const formData = new FormData();
    formData.append("data", text);
    setIsRequested(true);

    fetch(baseURL() + "/init", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        if (response.ok) {
          onInitServer(text);
        } else {
          console.log(response);
          setIsRequested(false);
        }
      })
      .catch((error) => {
        console.error(error);
        setIsRequested(false);
      });
  };

  return <Scanner onResult={(text, result) => handleScan(text)} />;
};

export default QRScanner;
