import prisma from '@/lib/db';
import type { Prisma, PromiseReturnType } from '@prisma/client';

export default async function getNavLinks() {
  const res = await prisma.category.findMany({
    orderBy: [
      {
        rank: 'asc',
      }
    ],
    include: {
      links: {
        orderBy: {
          rank: 'asc',
        },
        where: {
          public: true,
          status: 1,
        },
      },
    },
  });
  return res;
}

export type CategoryWithLinks = PromiseReturnType<typeof getNavLinks>