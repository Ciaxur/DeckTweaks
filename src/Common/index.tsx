import {
  gamepadDialogClasses,
  joinClassNames,
} from "decky-frontend-lib";

// Common Endpoints.
export const BACKEND_API = "http://localhost:3001";
export const BACKEND_WS = "ws://localhost:3001";

// Common Classes.
export const FieldWithSeparator = joinClassNames(gamepadDialogClasses.Field, gamepadDialogClasses.WithBottomSeparatorStandard);