import { useState } from "react";
import "./Upload.css";

async function upload(file: File) {
  const formData = new FormData();
  formData.set(file.name, file);
  const response = await fetch(
    `/api/upload/${file.name}`,
    {
      method: "PUT",
      body: formData,
    },
  );
  if (response.ok) {
    console.log(`upload ${file.name} successfully`);
  } else {
    console.log(`fail to upload ${file.name}`);
  }
}

export function Upload() {
  const [files, setFiles] = useState<FileList>();
  return (
    <form
      className="upload"
      onSubmit={async (e) => {
        e.preventDefault();
        if (!files) {
          return;
        }
        const arr = [];
        for (const f of files) {
          arr.push(upload(f));
        }
        await Promise.allSettled(arr);
      }}
    >
      <label>
        Choose Files
        <input
          type="file"
          multiple
          onChange={(e) => {
            if (e.target.files) {
              setFiles(e.target.files);
            }
          }}
        />
      </label>
      <button type="submit">submit</button>
    </form>
  );
}
