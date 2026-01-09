export default async function Itinerary({ params }: any) {
  const res = await fetch(process.env.NEXT_PUBLIC_API_BASE + "/api/itineraries");
  const data = await res.json();
  const item = data.find((x:any) => x.ID == params.id);

  async function buy() {
    "use server";
    const r = await fetch(process.env.NEXT_PUBLIC_API_BASE + "/api/create-order", {
      method: "POST",
      headers: { "Content-Type": "application/json" },
      body: JSON.stringify({ itineraryId: item.ID })
    });
    const o = await r.json();
    return o.orderId;
  }

  return (
    <main>
      <img src={item.ImageURL} width={600} />
      <h1>{item.Title}</h1>
      <p>{item.Description}</p>
      <p><b>Duration:</b> {item.Duration}</p>
      <p><b>Price:</b> ₹{item.Price}</p>
      <form action={buy}>
        <button>Buy & Download</button>
      </form>
    </main>
  );
}
