import { initializeApp } from "firebase/app";
import { getAuth, signInAnonymously } from "firebase/auth";

// get those from Firebase Project settings
const app = initializeApp({
  apiKey: "AIzaSyBi0OXLsAvPEXqgMLYvKox2UC8qa34cO-k",
  authDomain: "asoltd-guilds.firebaseapp.com",
  databaseURL: "https://asoltd-guilds-default-rtdb.firebaseio.com",
  projectId: "asoltd-guilds",
  storageBucket: "asoltd-guilds.appspot.com",
  messagingSenderId: "808202803134",
  appId: "1:808202803134:web:cc91ec3a36b7286e36cbef",
  measurementId: "G-B4H1DTT7Y6",
});

const auth = getAuth(app);

const credentials = await signInAnonymously(auth);
const idToken = await credentials.user.getIdToken(true);

console.log(idToken);
