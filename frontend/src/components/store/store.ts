// store.js
import { writable } from 'svelte/store';

// Create a writable store
const initialPath = 'Z:/'
export const path = writable(initialPath);

export const addEntryToPath = (newEntryName, oldPath: string) => {
    return oldPath + newEntryName + "/"
};

