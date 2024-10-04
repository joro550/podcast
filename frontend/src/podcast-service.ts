import { Weather } from "./podcast-models";

export default class PodcastService {
  private _baseUrl: string;

  constructor() {
    this._baseUrl = process.env.NEXT_PUBLIC_BACKEND_URL
      ?? '';
  }

  public async getWeather(): Promise<Weather[]> {
    let response = await fetch(`${this._baseUrl}/weather`);
    return (await response.json()) as Weather[];
  }
}
