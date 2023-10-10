import { uuidv4 } from "@firebase/util";
import { ApiClient } from ".";
import { initializeApp } from "firebase/app";
import { getAuth, signInWithEmailAndPassword } from "firebase/auth";

// get those from Firebase Project settings
const app = initializeApp({
  projectId: "",
  apiKey: "",
  authDomain: "",
  // ...
});

const auth = getAuth(app);

const email = "";
const password = "";
const credentials = await signInWithEmailAndPassword(auth, email, password);
const idToken = await credentials.user.getIdToken(true);

const client = new ApiClient({
  BASE: "https://api.lancr.co.uk",
  TOKEN: idToken, // getToken function that obtains firebase ID token
});

try {
  const res = await client.heroService.heroServiceCreateHero({
    payload: {
      id: uuidv4(),
    },
  });
  console.log(res);
} catch (err) {
  console.log(err);
}

try {
  const res = await client.heroService.heroServiceListHeroes();
  console.log(res);
} catch (err) {
  console.log(err);
}
