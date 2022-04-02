import { Header } from "../components/Header";
import { Upload } from "../components/Upload";
import { Pastebin } from "../components/Pastebin";
import "./App.css";

export function App() {
  return (
    <div className="App">
      <Header />
      <Upload />
      <Pastebin />
    </div>
  );
}
