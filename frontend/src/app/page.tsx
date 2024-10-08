"use client";

import Head from "next/head";
import "../../styles/globals.sass";
import Card from "./card";

export default function Home() {
  return (
    <div className="container">
      <Head>
        <title>Counter Strike Hot Takes</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="content">
        <h1>Counter Strike Hot Takes</h1>
      </main>

      <div className="fixed-grid has-3-cols">
        <div className="grid">
          <div className="cell">
            <Card
              src="/Kassad.jpg"
              alt="Image of Counter strike analyst and team owner Kassad"
            />
          </div>

          <div className="cell">
            <Card
              src="/Thorin.jpg"
              alt="Image of Counter strike analyst Thorin"
            />
          </div>

          <div className="cell">
            <Card
              src="/Mauisnake.jpg"
              alt="Image of Counter strike analyst mauisnake"
            />
          </div>
        </div>
      </div>
    </div>
  );
}
