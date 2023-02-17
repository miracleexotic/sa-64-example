import { PlaylistsInterface } from "./IPlaylist";
import { ResolutionsInterface } from "./IResolution";
import { VideosInterface } from "./IVideo";

export interface WatchVideoInterface {
  watched_time: Date,
  resolution_id: string,
  video_id: string,
}

export interface WatchVideoDataInterface {
  id: number,
  watched_time: Date,
  resolution_id: string,
  resolution: ResolutionsInterface,
  playlist_id: string,
  playlist: PlaylistsInterface,
  video_id: string,
  video: VideosInterface,
}
