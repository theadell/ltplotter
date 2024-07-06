import { DataSet } from "./models";

const apiBaseUrl = '/api';

export const Api = {
    async parseRawFile(file : File): Promise<DataSet> {
        const form = new FormData()
        form.append('file', file)
        const response = await fetch(`${apiBaseUrl}/parse`, {method: 'POST', body: form});
        if (!response.ok) {
          throw new Error('Network response was not ok');
        }
        const data: DataSet = await response.json();
        return data;
      },
}