import { FC, useEffect, useState } from "react";
import { BACKEND_WS, FieldWithSeparator } from "../Common";
import { Field } from "decky-frontend-lib";

interface Props {}

interface BatteryState {
  error:      string,
	currentNow: number,
	voltageNow: number,
	status:     string,
	capacity:   number,
}

export const BatteryState: FC<Props> = ({}: Props) => {
  const [state, setState] = useState<BatteryState>({
    error: "",
    capacity: 0,
    currentNow: 0,
    status: "Unknown",
    voltageNow: 0,
  });

  // Once the component is mounted, initialize websocket.
  useEffect(() => {
    // Establish a Websocket connection with the backend.
    console.log("BatteryState<Component>: Initializing telemetry websocket connection");
    const ws = new WebSocket(`${BACKEND_WS}/telemetry/battery`);
    ws.onmessage = msg => setState(JSON.parse(msg.data) as BatteryState);

    // Clean up on unmount.
    return () => {
      console.log('BatteryState<Component>: Tearing down');
      ws.close();
    }
  }, []);

  return (
    <div className={FieldWithSeparator}>
      <Field
        label="Capacity"
        description={state.capacity}
      />
      <Field
        label="Current"
        description={state.currentNow}
      />
      <Field
        label="Voltage"
        description={state.voltageNow}
      />
      <Field
        label="Charge State"
        description={state.status}
      />
      <Field
        label="Error"
        description={state.error}
      />
    </div>
  );
};