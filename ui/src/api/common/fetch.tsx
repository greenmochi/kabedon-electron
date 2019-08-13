import { MyAnimeListMsgType } from "../myanimelist/requestMessage";
import { AtlasMsgType } from "../atlas/requestMessage";

type msgType = 
 | MyAnimeListMsgType
 | AtlasMsgType
 ;

export function postFetch(endpoint: string, body: msgType): Promise<Response> {
  const options = {
    method: "POST",
    headers: {
      "Content-Type": "application/json",
    },
    body: JSON.stringify(body),
  };
  return fetch(endpoint, options)
}
