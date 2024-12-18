import Image from "next/image";
import { PieChart, Pie, ResponsiveContainer } from "recharts";

export type PresenterCard = {
  name: string;
  imageUrl: string;
  altText: string;
  username: string;
  description: string;
  socials: PresenterSocial[];
};

export type PresenterSocial = {
  username: string;
  url: string;
  icon: string;
};

export default function Card(presenter: PresenterCard) {
  const data01 = [
    { name: "Group A", value: 400 },
    { name: "Group B", value: 300 },
    { name: "Group C", value: 300 },
    { name: "Group D", value: 200 },
  ];
  return (
    <div className="card">
      <div className="card-image">
        <figure className="image is-4by3">
          <Image
            width={100}
            height={200}
            alt={presenter.altText}
            src={presenter.imageUrl}
          />
        </figure>
      </div>
      <div className="card-content">
        <div className="media">
          <div className="media-content">
            <p className="title is-4">{presenter.name}</p>
            {presenter.socials.map((s) => (
              <p className="subtitle is-6">
                <i className={s.icon}></i>
                <a href={s.url}>@{s.username}</a>
              </p>
            ))}
          </div>
        </div>
        <p>{presenter.description}</p>
        <ResponsiveContainer width="100%" height={200}>
          <PieChart width={400} height={500}>
            <Pie
              data={data01}
              dataKey="value"
              cx="50%"
              cy="50%"
              outerRadius={60}
              fill="#8884d8"
              label
            />
          </PieChart>
        </ResponsiveContainer>
      </div>
      <div className="card-footer">
        <a href="/presenter/" className="card-footer-item">
          More Information
        </a>
      </div>
    </div>
  );
}
