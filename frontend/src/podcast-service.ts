import { Weather } from "./podcast-models";

export default class PodcastService {
  private _baseUrl: string;

  constructor() {
    this._baseUrl = process.env.BACKEND_URL;
  }

  public async getWeather() {
    let response = await fetch(`${this._baseUrl}/weather`);
    return (await response.json()) as Weather[];
  }
}
