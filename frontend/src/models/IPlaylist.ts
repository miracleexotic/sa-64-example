import { WatchVideoInterface } from "./IWatchVideo";

export interface PlaylistsInterface {
    id: string,
    title: string,
    owner_id: number,
    watch_videos: WatchVideoInterface[],
  }
  