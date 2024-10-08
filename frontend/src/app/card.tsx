import Image from "next/image";
import { PieChart, Pie, ResponsiveContainer } from "recharts";

export default function Card({ src, alt }: { src: string; alt: string }) {
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
          <Image width={100} height={200} alt={alt} src={src} />
        </figure>
      </div>
      <div className="content">
        Lorem ipsum dolor sit amet, consectetur adipiscing elit. Phasellus nec
        iaculis mauris. <a>@bulmaio</a>. <a href="#">#css</a>
        <a href="#">#responsive</a>
        <br />
        <time>11:09 PM - 1 Jan 2016</time>
        <ResponsiveContainer width={100} height={100}>
          <PieChart width={400} height={400}>
            <Pie
              data={data01}
              dataKey="value"
              cx="50%"
              cy="50%"
              outerRadius={60}
              fill="#8884d8"
            />
          </PieChart>
        </ResponsiveContainer>
      </div>
    </div>
  );
}
