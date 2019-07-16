import { 
  compose,
  createStore,
  combineReducers,
  applyMiddleware,
} from "redux";
import thunk, { ThunkAction, ThunkMiddleware } from "redux-thunk";
import { devToolsEnhancer } from "redux-devtools-extension";

import { YoutubeDLActionType } from "./youtube-dl/type";
import { youtubeDLReducer } from "./youtube-dl/reducer";
import { NyaaActionType } from "./nyaa/type";
import { nyaaReducer } from "./nyaa/reducer";
import { CalculatorActionType } from "./calculator/type";
import { calculatorReducer } from "./calculator/reducer";
import { APIActionType } from "./api/type";
import { apiReducer } from "./api/reducer";
import { NavigationActionType } from "./navigation/type";
import { navigationReducer } from "./navigation/reducer";

const rootReducer = combineReducers({
  api: apiReducer,
  navigation: navigationReducer,
  nyaa: nyaaReducer,
  youtubeDL: youtubeDLReducer,
  calculator: calculatorReducer,
});

export type StoreState = ReturnType<typeof rootReducer>;
export type StoreActions =
  | APIActionType
  | NavigationActionType
  | NyaaActionType
  | YoutubeDLActionType
  | CalculatorActionType;
export type ThunkResult<R> = ThunkAction<R, StoreState, undefined, StoreActions>;

export default function configureStore() {
  if (process.env.NODE_ENV === "development") {
    return createStore(
      rootReducer, 
      compose(
        applyMiddleware(thunk as ThunkMiddleware<StoreState, StoreActions>), 
        devToolsEnhancer({}),
      ),
    );
  } else {
    return createStore(
      rootReducer, 
      applyMiddleware(thunk as ThunkMiddleware<StoreState, StoreActions>),
    );
  }
};