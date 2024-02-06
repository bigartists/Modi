import { baseQueryApi } from 'redux/services/baseQuery'
import { IJobsParams, IResultAppss } from '../types'
export const AppssApi = baseQueryApi.injectEndpoints({
  endpoints: builder => ({
    getAppss: builder.query<IResultAppss, { page?: number; pageSize?: number }>(
      {
        query: ({ page, pageSize }) => ({
          url: '/apps/api/',
          method: 'GET',
          params: {
            form_data: JSON.stringify({
              filters: [
                {
                  col: 'labels',
                  opr: 'ct',
                  value: '"kind": "job"',
                },
              ],
              page: page || 0,
              page_size: pageSize || 20,
            }),
          },
        }),
        transformResponse: (response: IResultAppss, meta, arg) => {
          return {
            data: response.data,
            count: response.count,
            label_columns: response.label_columns,
          }
        },

        providesTags: (result, error, arg) => {
          return result && result.data
            ? [
                { type: 'Post', id: 'LIST' },
                ...result?.data.map(({ id }) => ({
                  type: 'Post' as const,
                  id,
                })),
              ]
            : [{ type: 'Post', id: 'LIST' }]
        },
        // 该函数将在请求开始时调用，在处理程序中运行其他逻辑
        onQueryStarted: () => {},
      },
    ),

    addNewAppss: builder.mutation({
      query: initialPost => {
        return {
          url: '/apps/api/training',
          method: 'POST',
          body: initialPost,
        }
      },
      invalidatesTags: ['Post'],
    }),

    stopAppss: builder.mutation<any, { id: number }>({
      query: ({ id }) => {
        return {
          url: `/apps/api/stop/${id}`,
          method: 'POST',
        }
      },
      invalidatesTags: (result, error, arg) => [{ type: 'Post', id: arg.id }],
    }),

    getAppsParamsById: builder.query<IJobsParams, { id: number }>({
      query: ({ id }) => ({
        url: `apps/api/${id}`,
        method: 'GET',
      }),
    }),
  }),
})

export const {
  useGetAppssQuery,
  useLazyGetAppssQuery,
  useAddNewAppssMutation,
  useStopAppssMutation,
  useLazyGetAppsParamsByIdQuery,
  useGetAppsParamsByIdQuery,
} = AppssApi
