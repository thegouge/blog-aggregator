import Header from "@/components/Header";
import Footer from "@/components/Footer"

export default async function Home() {
  const feeds = await getAllFeeds()

  return (
    <div className="flex min-h-screen flex-col items-center  justify-between p-24">
      <Header />
      <main className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
        {feeds.map((feed) => {
          return <div key={feed.ID}>{feed.Name}</div>
        })}
      </main>

      <Footer />

    </div>
  );
}

async function getAllFeeds() {
  const response = await fetch("http://localhost:8000/v1/feeds")

  if (!response.ok) {
    throw new Error("failed to get all feeds")
  }

  return response.json()
}

