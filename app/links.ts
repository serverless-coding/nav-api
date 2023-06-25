import prisma from '@/lib/db';
import type { Prisma, Link as SiteLink, Category } from '@prisma/client';

// export default async function getNavLinks() {
//   const res = await prisma.category.findMany({
//     orderBy: [
//       {
//         rank: 'asc',
//       }
//     ],
//     include: {
//       links: {
//         orderBy: {
//           rank: 'asc',
//         },
//         where: {
//           public: true,
//           status: 1,
//         },
//       },
//     },
//   });
//   return res;
// }

export default async function getNavLinks(): Promise<CategoryWithLinks[]> {
  const res = await fetch("https://nav.programnotes.cn/api/link");
  const data: Data = await res.json();
  console.log("fetch category res:", data);
  data.data.map((c) => {
    if (c.links == undefined) {
      c.links=[];
    }
  })
  return data.data;
}

interface Data {
  code: number;
  message: string | null | undefined;
  data: CategoryWithLinks[];
}


export interface CategoryWithLinks {
  id: string;
  icon: string;
  title: string;
  description: string;
  rank: number;
  created_at: string;
  updated_at: string;
  links: SiteLink[];
}

// export type CategoryWithLinks = Prisma.PromiseReturnType<typeof getNavLinks>