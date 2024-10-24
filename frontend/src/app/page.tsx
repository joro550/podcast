"use client";

import Head from "next/head";
import "../../styles/globals.sass";
import Card from "./card";
import { useEffect, useState } from "react";
import { Presenter } from "../podcast-models";
import PodcastService from "../podcast-service";

export default function Home() {
  const [data, setData] = useState<Presenter[]>([]);

  useEffect(() => {
    const service = new PodcastService();
    service.getPresenters().then((presenters) => setData(presenters));
  }, []);

  return (
    <div className="container">
      <Head>
        <title>Counter Strike Hot Takes</title>
        <link rel="icon" href="/favicon.ico" />
      </Head>

      <main className="content pt-4">
        <h1>Counter Strike Hot Takes</h1>
        <hr />
      </main>

      <div className="fixed-grid has-3-cols">
        <div className="grid">
          {data.map((d) => (
            <div className="cell" key={d.id}>
              <Card
                imageUrl="/Kassad.jpg"
                name={d.name}
                altText="kassad"
                username="Kassad"
                description="this is a description"
              />
            </div>
          ))}

          {/* <div className="cell"> */}
          {/*   <Card */}
          {/*     src="/Thorin.jpg" */}
          {/*     alt="Image of Counter strike analyst Thorin" */}
          {/*   /> */}
          {/* </div> */}
          {/**/}
          {/* <div className="cell"> */}
          {/*   <Card */}
          {/*     src="/Mauisnake.jpg" */}
          {/*     alt="Image of Counter strike analyst mauisnake" */}
          {/*   /> */}
          {/* </div> */}
        </div>
      </div>
    </div>
  );
}
