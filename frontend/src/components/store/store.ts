// store.js
import { writable } from 'svelte/store';

// Create a writable store
const initialPath = 'Z:/'
export const path = writable(initialPath);

export const addEntryToPath = (newEntryName, oldPath: string) => {
    return oldPath + newEntryName + "/"
};


export const removeEntryFromPath = (oldPath: string) => {
    let thePath = oldPath.slice(0, -1);
    const lastSeparatorIndex = thePath.lastIndexOf('/');
    const separatorCount = thePath.split("/").length - 1;
    if (separatorCount <= 0) {
        return oldPath;
    }
    if (lastSeparatorIndex !== -1) {
        return oldPath.substring(0, lastSeparatorIndex) + "/";
    } else {
        // If there is no path separator, return the original path
        return oldPath;
    }
}
