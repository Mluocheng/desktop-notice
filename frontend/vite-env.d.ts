declare global {
    interface Window {
        DesktopNotice: {
            myGlobalFunction: () => viod,
            closeWindow: () => void;
            request: (url: string, options?: AxiosRequestConfig) => Promise<any>
        };
    }
  }
  
  export {};