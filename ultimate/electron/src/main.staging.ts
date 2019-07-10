import { app } from "electron";
import * as path from "path";

import { KokoroServer } from "./kokoroServer";
import { IpcMain } from "./ipcMain";
import { MainWindow } from "./browserWindow";

let mainWindow: MainWindow | null;
let kokoroServer: KokoroServer | null = null;
let ipcMain: IpcMain = new IpcMain();

app.on("ready", () => {
  let binaryPath: string = path.resolve("./service");
  if (binaryPath.length > 0) {
    kokoroServer = new KokoroServer("ultimate-kokoro.exe", binaryPath, "localhost", 9111, 9990);
    kokoroServer.run();

    console.log(`Running kokoro server. binaryPath=${binaryPath} endpoint=${kokoroServer.kokoroEndpoint}`);

    ipcMain.registerKokoroServerListener(kokoroServer);
  } else {
    console.log("binaryPath length is not valid: length=", binaryPath.length);
  }
  mainWindow = new MainWindow("http://localhost:3000", false, 900, 1200);
  mainWindow.create();
  mainWindow.registerWindowsButtonListener();
  mainWindow.registerDevtools();
  mainWindow.sendAfter("ultimate:kokoroServerEndpointResponse", kokoroServer.kokoroEndpoint);
  mainWindow.sendAfter("ultimate:gatewayServerEndpointResponse", kokoroServer.gatewayEndpoint);
});

app.on("window-all-closed", async () => {
  if (kokoroServer) {
    await kokoroServer.close();
  }

  if (process.platform !== "darwin") {
    app.quit();
  }
});

app.on("activate", () => {
  if (mainWindow === null) {
    mainWindow = new MainWindow("http://localhost:3000", false, 900, 1200);
    mainWindow.create();
    mainWindow.registerDevtools();
    mainWindow.sendAfter("kokoro-endpoint", kokoroServer.kokoroEndpoint);
    mainWindow.sendAfter("gateway-endpoint", kokoroServer.gatewayEndpoint);
  }
});