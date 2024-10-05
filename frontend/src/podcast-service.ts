"use client";
import { Weather } from "./podcast-models";

export default class PodcastService {
  private _baseUrl: string;

  constructor() {
    this._baseUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "";
    console.log("Getting backend url", process.env.NEXT_PUBLIC_BACKEND_URL);
    var env = process.env;

    Object.keys(env).forEach(function (key) {
      console.log("export " + key + '="' + env[key] + '"');
    });
  }

  public async getWeather(): Promise<Weather[]> {
    let response = await fetch(`${this._baseUrl}/weather`);
    return (await response.json()) as Weather[];
  }
}
