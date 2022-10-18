import {
  ButtonItem,
  definePlugin,
  DialogButton,
  gamepadDialogClasses,
  PanelSection,
  PanelSectionRow,
  Router,
  ServerAPI,
  staticClasses,
  joinClassNames,
  ToggleField,
  SliderField,
} from "decky-frontend-lib";
import React, { useEffect, useRef, useState, VFC } from 'react';
import {
  FaWhmcs,
  FaBreadSlice,
  FaHeart,
} from "react-icons/fa";

const BACKEND_API = "http://localhost:3001";

interface IBatteryMonitor {
  enabled:          boolean,
  max_charge_limit: number,
  min_charge_limit: number,
}

interface ISettingsResponse {
  message:  string,
  settings: {
    battery_monitor: IBatteryMonitor,
  },
};

interface ISetSettingsRequest {
  battery: IBatteryMonitor,
};

const Content: VFC<{ serverAPI: ServerAPI }> = ({ serverAPI }) => {
  const FieldWithSeparator = joinClassNames(gamepadDialogClasses.Field, gamepadDialogClasses.WithBottomSeparatorStandard);

  // Battery Monitor State.
  const [batteryMonitorState, setBatteryMonitorState] = useState<IBatteryMonitor>({
    enabled: false,
    min_charge_limit: 80,
    max_charge_limit: 30,
  })
  const batteryMonitorUpdateTimeoutId = useRef<NodeJS.Timeout|null>(null);
  const batteryMonitorUpdateCouter = useRef<number>(0);

  // Request the current status from the backend when mounting this Component.
  useEffect(() => {
    fetch(`${BACKEND_API}/status/settings`)
      .then(res => res.json())
      .then((res: ISettingsResponse) => setBatteryMonitorState(res.settings.battery_monitor))
      .catch(err => console.log(err));
  }, []);

  // Update the backend status after a no-change delay, so that the backend isn't spammed.
  useEffect(() => {
    batteryMonitorUpdateCouter.current++;
    console.log(`Battery Monitor: State changed[counter=${batteryMonitorUpdateCouter.current}] -> `, batteryMonitorState);

    // Only update after mounting (BUG: onMount useEffect is called twice).
    if (batteryMonitorUpdateCouter.current <= 2) return;

    // Reset timer.
    if (batteryMonitorUpdateTimeoutId.current !== null) {
      clearTimeout(batteryMonitorUpdateTimeoutId.current);
    }

    batteryMonitorUpdateTimeoutId.current = setTimeout(() => {
      // Reset timeout id.
      batteryMonitorUpdateTimeoutId.current = null;

      console.log("Battery Monitor: Updating Backend")
      fetch(`${BACKEND_API}/status/settings`, {
        method: 'POST',
        headers: {
          'Content-Type': 'application/json',
        },
        body: JSON.stringify({
          battery: batteryMonitorState,
        } as ISetSettingsRequest),
      })
        .then(res => res.json())
        .then(res => console.log('Battery Monitor: Backend updated -> ', res))
        .catch(err => console.log('Battery Monitor: Failed to updated backend -> ', err));
    }, 1000);
  }, [batteryMonitorState]);

  return (
    <PanelSection>
      {/* BATTERY HEALTH ROW */}
      <PanelSectionRow>
        <div className={FieldWithSeparator}>
          <div className={staticClasses.PanelSectionTitle}>Battery Monitor</div>
          <ToggleField
            label="Health Monitor"
            description="Toggle monitoring the battery's charge health."
            icon={<FaHeart/>}
            checked={batteryMonitorState.enabled}
            onChange={() => setBatteryMonitorState({
              ...batteryMonitorState,
              enabled: !batteryMonitorState.enabled,
            })}
          />

          {/* Battery Charge Limits. */}
          {batteryMonitorState.enabled && (
            <React.Fragment>
              <SliderField
                label="Maximum charge limit"
                value={batteryMonitorState.max_charge_limit}
                showValue={true}
                onChange={value => setBatteryMonitorState({
                  ...batteryMonitorState,
                  max_charge_limit: value,
                })}
                min={0}
                max={100}
              />
              <SliderField
                label="Minimum charge limit"
                value={batteryMonitorState.min_charge_limit}
                showValue={true}
                onChange={value => setBatteryMonitorState({
                  ...batteryMonitorState,
                  min_charge_limit: value,
                })}
                min={0}
                max={100}
              />
            </React.Fragment>
          )}
        </div>
      </PanelSectionRow>

      {/* DEBUG ROW */}
      <PanelSectionRow>
        <div className={FieldWithSeparator}>
          <div className={staticClasses.PanelSectionTitle}>
            Debug
          </div>
          {/* Toast Testing. */}
          <ButtonItem
              layout="below"
              onClick={() => serverAPI.toaster.toast({
                title: "I am the toast!",
                body: "The toast shalt be mine, 🍞",
                icon: <FaBreadSlice/>,
              })}
            >
            ToastTest
          </ButtonItem>
        </div>
      </PanelSectionRow>
    </PanelSection>
  );
};


const DeckyPluginRouterTest: VFC = () => {
  return (
    <div style={{ marginTop: "50px", color: "white" }}>
      Hello World!
      <DialogButton onClick={() => Router.NavigateToStore()}>
        Go to Store
      </DialogButton>
    </div>
  );
};

export default definePlugin((serverApi: ServerAPI) => {
  serverApi.routerHook.addRoute("/decky-plugin-test", DeckyPluginRouterTest, {
    exact: true,
  });

  return {
    title: <div className={staticClasses.Title}>DeckTweaks</div>,
    content: <Content serverAPI={serverApi} />,
    icon: <FaWhmcs />,
    onDismount() {
      serverApi.routerHook.removeRoute("/decky-plugin-test");
    },
  };
});
