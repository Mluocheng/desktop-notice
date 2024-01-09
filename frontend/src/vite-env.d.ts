/// <reference types="vite/client" />

export { }
interface Runtime {
    WindowSetPosition: Function,
    WindowSetAlwaysOnTop: Function,
    Hide: Function,
    Show: Function,
    WindowSetMaxSize: Function
    
}
declare global {
  interface Window { runtime: Runtime; }
}
