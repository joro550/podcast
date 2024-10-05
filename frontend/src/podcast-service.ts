import { Weather } from "./podcast-models";

export default class PodcastService {
  private _baseUrl: string;

  constructor() {
    this._baseUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "";
    console.log("Getting backend url", process.env.NEXT_PUBLIC_BACKEND_URL);
    console.log("ENV", process.env);
  }

  public async getWeather(): Promise<Weather[]> {
    let response = await fetch(`${this._baseUrl}/weather`);
    return (await response.json()) as Weather[];
  }
}
