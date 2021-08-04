import Axios from 'axios';

const API_URL = process.env.VUE_APP_API_URL || `${window.location.origin}/api`;

export const RestApi = {
  work: {
    async getList() {
      return (await Axios.get(`${API_URL}/work/list`)).data.response;
    },
    async get(id: string) {
      return (await Axios.get(`${API_URL}/work?id=${id}`)).data.response;
    },
    async update(todo: any) {
      return (await Axios.patch(`${API_URL}/work`, todo)).data.response;
    },
    async add(title: string, tags: string[]) {
      return (
        await Axios.post(`${API_URL}/work`, {
          title,
          tags,
        })
      ).data.response;
    },
    async delete(id: string) {
      return (await Axios.delete(`${API_URL}/work?id=${id}`)).data.response;
    },
  },

  image: {
    async update(image: any) {
      return (await Axios.patch(`${API_URL}/image`, image)).data.response;
    },
    async add(workId: string, time: number, image: string, created: string) {
      return (
        await Axios.post(`${API_URL}/image`, {
          workId,
          time,
          image,
          created,
        })
      ).data.response;
    },
    async delete(workId: string, id: string) {
      return (await Axios.delete(`${API_URL}/image?workId=${workId}&id=${id}`)).data.response;
    },
  },
};
