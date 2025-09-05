// import {
//   Table,
//   TableBody,
//   TableCell,
//   TableHead,
//   TableHeader,
//   TableRow,
// } from './components/ui/table';

import { Card, CardContent, CardHeader, CardTitle } from '@/components/ui/card';

// import { Badge } from './components/ui/badge';

import {
  ArrowRight,
  Warehouse,
  Construction,
  Package,
  Calendar,
  MapPin,
  Map,
} from 'lucide-react';

import { format } from 'date-fns';


import { ADMIN_ID } from '@/models/auth';
import type { CustomLog } from '@/models/customLog';
import { FetchAllLogs } from '@/service/auth.api';
import React, { useMemo, useEffect, useState } from 'react'
import { Table, TableBody, TableCell, TableHead, TableHeader, TableRow } from '@/components/ui/table';
import { Badge } from '@/components/ui/badge';

const ReportPage = () => {
  const [customLogs, setCustomLogs] = useState<CustomLog[]>([]);


  useEffect(()=>{

    const fetchLogsByAdmin = async() => {
      const adminId = localStorage.getItem(ADMIN_ID)!;

      if (adminId) {
        const res = await FetchAllLogs(adminId);

        if (res.success) {
          setCustomLogs(res.data);
        }

      } else {
        console.log("No admin id in logs page");
      }
    }

    fetchLogsByAdmin()
  },[]);

  const formatAddress  = (location:string, state : string, country: string) => {
      return `${location}, ${state}, ${country}`;
  };

  return (
    <div
      className='min-h-screen sm:p-6 p-4 flex  justify-center '
    >
        <div className='w-full max-w-7xl'>
          <Card>
              <CardHeader>
                <CardTitle>Construction Logistics Logs</CardTitle>
              </CardHeader>

              <CardContent>
                <Table>
                  <TableHeader>
                    <TableRow>
                
                      <TableHead className='min-w-[150px]'>From</TableHead>
                      <TableHead className="min-w-[50px] text-center">Movement</TableHead>
                      <TableHead className="min-w-[150px]">To</TableHead>
                      <TableHead className="min-w-[150px]">Material</TableHead>
                      <TableHead className="min-w-[100px] text-right">Quantity</TableHead>
                      <TableHead className="min-w-[150px] text-right">Updated At</TableHead>
                    
                    </TableRow>
                  </TableHeader>

                  <TableBody>
                    {customLogs.map((log, index) => (
                      <TableRow key={index}>
                          <TableCell>
                            <div className='flex flex-col space-y-1'>
                                <div className='flex items-center gap-2'>
                                    <Warehouse className='w-4 h-4 text-primary'/>
                                    <span className='font-medium'>{log.from_warehouse_name}</span>
                                </div>

                                <div
                                  className='flex items-center gap-2 text-muted-foreground text-sm'
                                >
                                  <MapPin className='w-4 h-4'/>
                                  <span>{formatAddress(log.from_warehouse_location, log.from_warehouse_state, log.from_warehouse_country)}
                                  </span>
                                </div>
                            </div>
                          </TableCell>

                          <TableCell
                            className='text-center'
                          >
                              <div
                                className='flex items-center justify-center gap-2'
                              > 
                                      <Badge className={`
                                            ${log.is_site ? 'border-amber-500 bg-amber-500/10 text-amber-500 hover:bg-amber-500/20' : 'border-blue-500 bg-blue-500/10 text-blue-500 hover:bg-blue-500/20'}
                                        `}>
                                          {log.is_site ? 'Site' : 'Warehouse'}
                                      </Badge>
                                      <ArrowRight className="w-4 h-4 text-muted-foreground" />
                                  
                              </div>
                          </TableCell>

                          <TableCell>
                            <div className="flex flex-col space-y-1">
                              <div className="flex items-center gap-2">
                                {log.is_site ? <Construction className="w-4 h-4 text-primary" /> : <Warehouse className="w-4 h-4 text-primary" />}
                                <span className="font-medium">{log.to_destination_name}</span>
                              </div>
                              <div className="flex items-center gap-2 text-muted-foreground text-sm">
                                <Map className="w-4 h-4" />
                                <span>{formatAddress(log.to_destination_location, log.to_destination_state, log.to_destination_country)}</span>
                              </div>
                            </div>
                          </TableCell>

                          <TableCell>
                            <div className="flex items-center gap-2 text-foreground">
                              <Package className="w-4 h-4" />
                              <span className="font-medium">{log.supply_name}</span>
                            </div>
                          </TableCell>
                         
                          <TableCell className="text-right">
                            <span className="font-medium">{log.supply_quantity}</span>
                            <span className="text-muted-foreground text-sm ml-1">{log.supply_unit}</span>
                          </TableCell>

                          <TableCell className="text-right">
                            <div className="flex items-center justify-end gap-2 text-muted-foreground text-sm">
                              <Calendar className="w-4 h-4" />
                              <span>{format(new Date(log.updated_time), 'PPp')}</span>
                            </div>
                          </TableCell>
                      </TableRow>
                    ))}
                  </TableBody>
                </Table>
              </CardContent>
          </Card>
        </div>
    </div>
  )
}

export default ReportPage