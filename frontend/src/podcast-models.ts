"use client";

export type Presenter = {
  name: string;
  description: string;
  id: number;
  imageUrl: string;
  altText: string;
  socials: Social[];
};

export type Social = {
  username: string;
  url: string;
  icon: string;
};
