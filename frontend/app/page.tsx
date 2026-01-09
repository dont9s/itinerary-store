export default async function Home() {
  let data: any[] = [];

  try {
    const res = await fetch(
      `${process.env.NEXT_PUBLIC_API_BASE}/api/itineraries`,
      { cache: "no-store" }
    );

    if (!res.ok) {
      console.error("API failed", res.status);
    } else {
      const json = await res.json();
      if (Array.isArray(json)) {
        data = json;
      }
    }
  } catch (err) {
    console.error("Fetch error", err);
  }

  return (
    <main>
      <h1>Travel Itineraries</h1>

      {data.length === 0 && (
        <p>No itineraries available.</p>
      )}

      <div
        style={{
          display: "grid",
          gridTemplateColumns: "repeat(3, 1fr)",
          gap: 20,
        }}
      >
        {data.map((i: any) => (
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
