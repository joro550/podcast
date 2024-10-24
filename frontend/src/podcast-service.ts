"use client";

import { Presenter } from "./podcast-models";

export default class PodcastService {
  private _baseUrl: string;

  constructor() {
    this._baseUrl = process.env.NEXT_PUBLIC_BACKEND_URL ?? "";
  }

  public async getPresenters(): Promise<Presenter[]> {
    let response = await fetch(`${this._baseUrl}/presenters`);
    return (await response.json()) as Presenter[];
  }
}
