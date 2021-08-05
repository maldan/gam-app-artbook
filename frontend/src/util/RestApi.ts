import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  work: {
    async getList(): Promise<any> {
      return (await Axios.get(`${API_URL}/work/list`)).data.response;
    },
    async get(id: string): Promise<any> {
      return (await Axios.get(`${API_URL}/work?id=${id}`)).data.response;
    },
    async update(todo: any): Promise<any> {
      return (await Axios.patch(`${API_URL}/work`, todo)).data.response;
    },
    async add(title: string, tags: string[]): Promise<any> {
      return (
        await Axios.post(`${API_URL}/work`, {
          title,
          tags,
        })
      ).data.response;
    },
    async delete(id: string): Promise<any> {
      return (await Axios.delete(`${API_URL}/work?id=${id}`)).data.response;
    },
  },

  image: {
    async get(workId: string, id: string): Promise<any> {
      return (await Axios.get(`${API_URL}/image?workId=${workId}&id=${id}`)).data.response;
    },
    async update(image: any): Promise<any> {
      return (await Axios.patch(`${API_URL}/image`, image)).data.response;
    },
    async moveDown(workId: string, id: string): Promise<any> {
      return (
        await Axios.patch(`${API_URL}/image/moveDown`, {
          id,
          workId,
        })
      ).data.response;
    },
    async add(workId: string, time: number, image: Blob, created: string): Promise<any> {
      const formData = new FormData();
      formData.append('workId', workId);
      formData.append('time', time + '');
      formData.append('image', image);
      formData.append('created', created);

      return (await Axios.post(`${API_URL}/image`, formData)).data.response;
    },
    async delete(workId: string, id: string): Promise<any> {
      return (await Axios.delete(`${API_URL}/image?workId=${workId}&id=${id}`)).data.response;
    },
  },
};
