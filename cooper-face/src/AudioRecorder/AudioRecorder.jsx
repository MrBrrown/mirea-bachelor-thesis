import * as React from "react";
import { AudioRecorder } from "react-audio-voice-recorder";
import baseURL from "../../config";

const AudioRecorderComponent = ({ connectionString, onServerResponse }) => {
  const addAudioElement = (blob) => {
    const reader = new FileReader();
    reader.readAsDataURL(blob);
    reader.onloadend = () => {
      const formData = new FormData();
      const base64data = reader.result;
      const base64 = base64data.split(",")[1];
      formData.append("data", base64);
      formData.append("auth", connectionString);

      fetch(baseURL() + "/command", {
        method: "POST",
        body: formData,
      })
        .then((response) => {
          if (response.ok) {
            response.json().then((data) => {
              onServerResponse(data.text);
            });
          } else {
            onServerResponse("Audio upload failed" + response.data);
          }
        })
        .catch((error) => onServerResponse("Error uploading audio:" + error));
    };
  };

  return (
    <div>
      <AudioRecorder
        onRecordingComplete={addAudioElement}
        audioTrackConstraints={{
          noiseSuppression: true,
          echoCancellation: true,
          // autoGainControl,
          // channelCount,
          // deviceId,
          // groupId,
          // sampleRate,
          // sampleSize,
        }}
        onNotAllowedOrFound={(err) => console.table(err)}
        downloadOnSavePress={false}
        downloadFileExtension="webm"
        mediaRecorderOptions={{
          audioBitsPerSecond: 128000,
        }}
        // showVisualizer={true}
      />
      <br />
    </div>
  );
};

export default AudioRecorderComponent;
