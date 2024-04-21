import { AuthActions, LocalStorageElements } from "@/shared/enums";
import { IUser } from "@/shared/interfaces";
import { Reducer } from "react";

export interface AuthAction {
  type: AuthActions;
  // eslint-disable-next-line @typescript-eslint/no-explicit-any
  payload?: any;
}

export const AuthReducer: Reducer<IUser | undefined, AuthAction> = (
  state,
  action
) => {
  switch (action.type) {
    case AuthActions.SetUser: {
      const user = JSON.stringify(action.payload);
      window.localStorage.setItem(LocalStorageElements.User, user);
      return action.payload;
    }
    case AuthActions.ClearUser: {
      window.localStorage.removeItem(LocalStorageElements.User);
      return null;
    }
    case AuthActions.RefreshUser: {
      const user = window.localStorage.getItem(LocalStorageElements.User);
      if (!user) {
        return null;
      }
      return JSON.parse(user);
    }
    default:
      return state;
  }
};