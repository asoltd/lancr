import { HeroServiceClient } from ".";

const client = new HeroServiceClient({
  BASE: "https://api.lancr.co.uk",
  TOKEN: "", // getToken function that obtains firebase ID token
});

try {
  const res = await client.heroService.heroServiceListHeroes();
  console.log(res);
} catch (err) {
  console.log(err);
}
