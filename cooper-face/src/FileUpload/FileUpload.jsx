import React, { useState } from "react";

import getAPI from "../../config";

export default function FileUploader() {
  const [image, setImage] = useState(null);

  const handleImageChange = (event) => {
    setImage(event.target.files[0]);
  };

  const handleImageUpload = (event) => {
    event.preventDefault();

    const formData = new FormData();
    formData.append("image", image);

    fetch(getAPI() + "/uploadFile", {
      method: "POST",
      body: formData,
    })
      .then((response) => {
        console.log(response);
      })
      .catch((error) => {
        console.error(error);
      });
  };

  return (
    <div>
      <input type="file" onChange={handleImageChange} />
      <button onClick={handleImageUpload}>Upload</button>
    </div>
  );
}
