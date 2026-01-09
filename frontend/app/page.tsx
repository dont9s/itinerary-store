export default async function Home() {
  const res = await fetch(process.env.NEXT_PUBLIC_API_BASE + "/api/itineraries", { cache: "no-store" });
  const data = await res.json();

  return (
    <main>
      <h1>Travel Itineraries</h1>
      <div style={{display:'grid',gridTemplateColumns:'repeat(3,1fr)',gap:20}}>
        {data.map((i:any) => (
          <a key={i.ID} href={`/itinerary/${i.ID}`}>
            <img src={i.ImageURL} width={300} />
            <h3>{i.Title}</h3>
            <p>{i.Duration}</p>
          </a>
        ))}
      </div>
    </main>
  );
}
