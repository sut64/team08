import React, { useEffect, useState } from "react";
import { Link as RouterLink } from "react-router-dom";
import { createStyles, makeStyles, Theme } from "@material-ui/core/styles";
import Typography from "@material-ui/core/Typography";
import Button from "@material-ui/core/Button";
import Container from "@material-ui/core/Container";
import Paper from "@material-ui/core/Paper";
import Box from "@material-ui/core/Box";
import Table from "@material-ui/core/Table";
import TableBody from "@material-ui/core/TableBody";
import TableCell from "@material-ui/core/TableCell";
import TableContainer from "@material-ui/core/TableContainer";
import TableHead from "@material-ui/core/TableHead";
import TableRow from "@material-ui/core/TableRow";
import { AmbulanceChecksInterface } from "../models/IAmbulanceCheck";
import { format } from 'date-fns'

const useStyles = makeStyles((theme: Theme) =>
  createStyles({
    container: {
      marginTop: theme.spacing(2),
    },
    table: {
      minWidth: 650,
    },
    tableSpace: {
      marginTop: 20,
    },
  })
);

function AmbulanceChecks() {
  const classes = useStyles();
  const [ambulancechecks, setAmbulanceCheck] = React.useState<AmbulanceChecksInterface[]>([]);
  const apiUrl = "http://localhost:8080/ambulancechecks"
  const requestOptions = {
    method: "GET",
    headers: { 
      Authorization: `Bearer ${localStorage.getItem("token")}`,
      "Content-Type": "application/json" },
  };
  
  const getAmbulanceChecks = async () => {
    fetch(`${apiUrl}`, requestOptions)
      .then((response) => response.json())
      .then((res) => {
        console.log(res.data);
        if (res.data) {
          setAmbulanceCheck(res.data);
        } else {
          console.log("else");
        }
      });
    };

  useEffect(() => {
    getAmbulanceChecks();
  }, []);

  return (
    <div>
      <Container className={classes.container} maxWidth="md">
        <Box display="flex">
          <Box flexGrow={1}>
            <Typography
              component="h2"
              variant="h6"
              color="primary"
              gutterBottom
            >
              บันทึกข้อมูล
            </Typography>
          </Box>
          <Box>
            <Button
              component={RouterLink}
              to="/ambulancecheck/create"
              variant="contained"
              color="primary"
            >
              สร้างข้อมูล
            </Button>
          </Box>
        </Box>
        <TableContainer component={Paper} className={classes.tableSpace}>
          <Table className={classes.table} aria-label="simple table">
            <TableHead>
              <TableRow>
                <TableCell align="center" width="10%">
                  รหัส
                </TableCell>
                <TableCell align="center" width="10%">
                  พนักงาน
                </TableCell>
                <TableCell align="center" width="10%">
                  รถ
                </TableCell>
                <TableCell align="center" width="30%">
                  ปัญหา
                </TableCell>
                <TableCell align="center" width="10%">
                  รหัสเอกสาร
                </TableCell>
                <TableCell align="center" width="10%">
                  ความรุนแรง
                </TableCell>
                <TableCell align="center" width="30%">
                  หมายเหตุ
                </TableCell>
                <TableCell align="center" width="30%">
                  วันที่และเวลา
                </TableCell>
              </TableRow>
            </TableHead>
            <TableBody>
              {ambulancechecks.map((item: AmbulanceChecksInterface) => (
                <TableRow key={item.ID}>
                  <TableCell align="center">{item.ID}</TableCell>
                  <TableCell align="center">{item.Recorder.ID}</TableCell>
                  <TableCell align="center">{item.Ambulance.ID}</TableCell>
                  <TableCell align="center">{item.Problem.ID}</TableCell>
                  <TableCell align="center">{item.DocCode}</TableCell>
                  <TableCell align="center">{item.Severity}</TableCell>
                  <TableCell align="center">{item.Note}</TableCell>
                  <TableCell align="center">{format((new Date(item.DateAndTime)), 'dd MMMM yyyy hh:mm')}</TableCell>
                </TableRow>
              ))}
            </TableBody>
          </Table>
        </TableContainer>
      </Container>
    </div>
  );
}

export default AmbulanceChecks;