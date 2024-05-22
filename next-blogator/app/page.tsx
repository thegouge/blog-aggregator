import Header from "@/components/Header";
import Footer from "@/components/Footer"

export default function Home() {
  return (
    <div className="flex min-h-screen flex-col items-center justify-between p-24">
      <Header />
      <main className="z-10 w-full max-w-5xl items-center justify-between font-mono text-sm lg:flex">
        This is the Main!
      </main>

      <Footer />

    </div>
  );
}
